requirejs.config({
  baseUrl: '..',
  paths: {
    axios:   '/js/lib/goa/axios.min',
    client:  '/js/lib/goa/client',
    ractive: '/js/lib/ractive/ractive',
    text:    '/js/lib/ractive/text',
    rv:      '/js/lib/ractive/rv',
  }
});

// Post List
requirejs([
  'client',
  'ractive',
  'rv!tmpl/post',
  'rv!tmpl/post_list',
  'rv!tmpl/post_form',
  'rv!tmpl/api_status',
], function (client, Ractive, post_tmpl, post_list_tmpl, post_form_tmpl, api_status_tmpl) {

  var HixIO = {};

  HixIO.PostItem = Ractive.extend({
    template: post_tmpl,
    onrender: function() {
      this.on({
        delete: function() {
          client().deletePost('/api/v1/posts/' + this.get('id'))
            .then(this.teardown.bind(this))
            .catch(function() {
              console.log('post deletion failed for id: ' + this.get('id'));
            });
        }
      });
    },
  });

  HixIO.PostList = Ractive.extend({
    template: post_list_tmpl,
    components: {
      post: HixIO.PostItem,
    },
    data: function() {
      return {
        posts: [],
        message: '',
      };
    },
    fetch: function() {
      var self = this;

      client().listPost('/api/v1/posts').then(function(response) {
        self.set('posts', response.data);
      }).catch(function(error) {
        self.set('message', error);
      });
    },
    onrender: function() {
      this.on('fetch', this.fetch);
      this.fetch(); 
    },
  });

  HixIO.APIStatus = Ractive.extend({
    template: api_status_tmpl,
    data: function() {
      return {
        status: "",
        status_class: ""
      };
    },
    check: function() {
      var self = this;
      client().checkStatus('/api/v1/status')
        .then(function(response) {
          self.set({
            status: response.statusText,
            status_class: 'status_ok',
          });
        }).catch(function(error) {
          self.set({
            status: "Not OK",
            status_class: 'status_error',
          });
        });
    },
    onrender: function() {
      this.on('check', this.check);
      this.check();
    },
  });

  HixIO.PostForm = Ractive.extend({
    template: post_form_tmpl,
    data: function() {
      return {
        post: {
          title: '',
          body: '',
          published: false,
        },
        message: '',
      };
    },
    onrender: function() {
      this.on({
        create: function() {
          var self;

          self = this;
          client().createPost('/api/v1/posts', this.get('post'))
            .then(function(response) {
              console.log('create post: ' + response);
              self.fire('success');
            })
            .catch(function(error) {
              self.set('message', error);
              self.fire('error');
            });
        }
      });
    },
  });

  var post_list = new HixIO.PostList({ el: '#post_list' });
  var post_form = new HixIO.PostForm({ el: '#post_form' });
  var api_status = new HixIO.APIStatus({ el: '#api_status' });

  post_form.on({
    success: function() {
      console.log('firing fetch');
      post_list.fire('fetch');
    }
  });

});

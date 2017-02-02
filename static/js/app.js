requirejs.config({
  baseUrl: '..',
  paths: {
    axios:      '/js/lib/goa/axios.min',
    client:     '/js/lib/goa/client',
    ractive:    '/js/lib/ractive/ractive',
    text:       '/js/lib/ractive/text',
    rv:         '/js/lib/ractive/rv',
    jwt_decode: '/js/lib/jwt-decode',
  }
});

requirejs([
  'client',
  'ractive',
  'jwt_decode',
  'rv!tmpl/blog',
  'rv!tmpl/post',
  'rv!tmpl/post_list',
  'rv!tmpl/post_form',
  'rv!tmpl/api_status',
  'rv!tmpl/auth_form',
], function (client, Ractive, jwt_decode, blog_tmpl, post_tmpl, post_list_tmpl, post_form_tmpl, api_status_tmpl, auth_form_tmpl) {

  /*
   * Date shim
   */
  Date.prototype.getUnixTime = function() { return this.getTime()/1000|0 };
  if(!Date.now) Date.now = function() { return new Date(); }
  Date.time = function() { return Date.now().getUnixTime(); }

  // Namespace
  var HixIO = {};

  HixIO.PostItem = Ractive.extend({
    template: post_tmpl,
    data: function() {
      return {
        message: null,
        message_class: null,
      };
    },
    onrender: function() {
      this.on({
        delete: function() {
          var self;

          self = this;

          opts = {
            headers: {}
          }
          if(this.get('auth')) {
            opts['headers']['authorization'] = this.get('auth.header');
          }
          client().deletePost('/api/v1/posts/' + self.get('id'), opts)
            .then(self.teardown.bind(self))
            .catch(function() {
              self.set({
                message: 'Failed to delete post.',
                message_class: 'error',
              });
            });
        },
        edit: function() {
          console.log('edit');
        },
        update_post: function() {
          console.log('Updating post: ', this.get());
          opts = {
            headers: {}
          }
          if(this.get('auth')) {
            opts['headers']['authorization'] = this.get('auth.header');
          }
          client().updatePost('/api/v1/posts', this.get('post'), opts)
            .then(function(response) {
              self.set({
                message: 'success',
                message_class: 'success',
              });
            })
            .catch(function(error) {
              self.set({
                message: error.statusText,
                message_class: 'error',
              });
            });
        },
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
        posts: null,
        message: null,
        message_class: null,
      };
    },
    fetch: function() {
      var self = this;

      client().listPost('/api/v1/posts').then(function(response) {
        self.set('posts', []); // this was apparently necessary, deleting a post then adding one didn't update the list.
        self.set('posts', response.data);
      }).catch(function(error) {
        self.set({
          message: 'Unable to fetch posts.',
          message_class: 'error',
        });
      });
    },
    onrender: function() {
      this.on('fetch', function() { this.fetch(); });
      this.fetch();
    },
  });

  HixIO.APIStatus = Ractive.extend({
    template: api_status_tmpl,
    data: function() {
      return {
        status: null,
        status_class: null,
      };
    },
    check: function() {
      var self = this;

      this.set({
        status: 'checking...',
        status_class: 'warning',
      });

      window.setTimeout(function() {
        client().checkStatus('/api/v1/status')
          .then(function(response) {
            self.set({
              status: response.statusText,
              status_class: 'success',
            });
          }).catch(function(error) {
            self.set({
              status: "Not OK",
              status_class: 'error',
            });
          });
      }, 1000);
    },
    onrender: function() {
      this.on('check', function() { this.check(); });
      this.check();
    },
  });

  HixIO.PostForm = Ractive.extend({
    template: post_form_tmpl,
    data: function() {
      return {
        post: {
          title: null,
          body: null,
          published: false,
        },
        message: null,
      };
    },
    onrender: function() {
      this.on({
        create: function() {
          var self;

          self = this;
          opts = {
            headers: {}
          }
          if(this.get('auth')) {
            opts['headers']['authorization'] = this.get('auth.header');
          }
          client().createPost('/api/v1/posts', this.get('post'), opts)
            .then(function(response) {
              self.set({
                message: 'Success!',
                message_class: 'success',
              });
              self.fire('success');
            })
            .catch(function(error) {
              self.set({
                message: 'Failed to create post.',
                message_class: 'error',
              });
            });
        }
      });
    },
  });

  HixIO.AuthForm = Ractive.extend({
    template: auth_form_tmpl,
    data: function() {
      return {
        user: {
          email: null,
          password: null,
        },
        message: null,
      };
    },
    onrender: function() {
      this.on({
        sign_in: function() {
          var self;

          self = this;

          client().jwtAuth('/api/v1/auth', this.get('user'))
            .then(function(response) {
              var auth_header, store;

              auth_header = response.headers.authorization;
              if(!auth_header) {
                self.set({
                  message: 'Unable to log in.',
                  message_class: 'error',
                });
                return;
              }

              self.get('auth_store').setItem(self.get('auth_key'), auth_header.split(/\s+/)[1]);
              self.set({
                user: {
                  email: null,
                  password: null,
                },
              });
              self.fire('signed-in');
            })
            .catch(function(error) {
              self.set({
                message: 'That didn\'t work.',
                message_class: 'error',
              });
            });
        },
        sign_out: function() {
          this.get('auth_store').removeItem(this.get('auth_key'));
          this.set({
            user: {
              email: null,
              password: null,
            },
          });
          this.fire('signed-out');
        },
      });
    },
  });

  var blog = new Ractive({
    el: '#blog',
    template: blog_tmpl,
    components: {
      apistatus: HixIO.APIStatus,
      authform: HixIO.AuthForm,
      postform: HixIO.PostForm,
      postlist: HixIO.PostList,
    },
    data: function() {
      return {
        auth_store: localStorage,
        auth_key: 'hixio.auth.token',
        auth: null,
      };
    },
    auth_sync: function() {
      var token;

      console.log('auth_sync');
      token = this.get('auth_store').getItem(this.get('auth_key'));

      if(token === null) {
        this.set('auth', null);
        return;
      }

      this.set({
        auth: {
          token: jwt_decode(token),
          header: 'Bearer ' + token,
        }
      });
    },
    oninit: function() {
      this.auth_sync();
    },
    onrender: function() {
      var self;

      self = this;

      this.on('*.signed-in *.signed-out', function() {
        self.auth_sync();
      });

      this.on('*.success', function() {
        var comps = self.findAllComponents('postlist');
        console.log(this.event);
        for(c in comps) {
          comps[c].fire('fetch');
        }
      });

      this.on('* *.*', function() {
        //console.log(this.event);
      });

    },
  });

});


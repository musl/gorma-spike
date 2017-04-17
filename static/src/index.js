import React, {
  Component
} from 'react';
import ReactDOM from 'react-dom';
import {
  HashRouter as Router,
  Link,
  Route,
  Switch
} from 'react-router-dom';
import Moment from 'react-moment';
import moment from 'moment';
import axios from 'axios';

const NotFound = (props) => (
  <div>
    <h2>{props.message ? props.mssage : 'Uh oh.'}</h2>
    <p>I couldn't find that{props.thing ? ` ${props.thing}` : null}.</p>
  </div>
);

const EndMark = () => (
  <span className="endmark">&nbsp;&#x221e;</span>
);

const Photo = (props) => (
  <div>
    <h2>{props.alt}</h2>
    <p><Moment format="YYYY-MM-DD">{props.created_at}</Moment></p>
    <a href={props.original}>
      <img src={props.original} alt={props.alt} className="pure-img photo-original"/>
    </a>
  </div>
);

const Thumbnail = (props) => (
  <Link to={"/photos/" + props.id} className="photo-thumbnail">
    <img src={props.thumbnail} alt={props.alt} className="photo-thumbnail"/>
  </Link>
);

const PhotoIndex = (props) => (
  <div>
    <h2>Photos</h2>
    {Object.keys(props.photo_groups).length === 0 ? (
      <p>please upload some photos.</p>
    ) : (
      Object.keys(props.photo_groups).map((group) => (
        <div key={group}>
          <h2>{group}</h2>
          {props.photo_groups[group].map((photo) => (
            <Thumbnail key={photo.id.toString()} {...photo}/>
          ))}
        </div>
      ))
    )}
  </div>
);

const Post = (props) => (
  <div>
    <h2><Link to={"/posts/" + props.id}>{props.title}</Link></h2>
    <p>
      <span dangerouslySetInnerHTML={{__html: props.body}} />
      <EndMark/>
    </p>
  </div>
);

const RecentPosts = (props) => (
  <div>
    <h2>Recent Posts</h2>
    <ul className="post-list">
      {props.posts.slice(0,5).map((post) =>
        <li key={post.id.toString()}>
          <Moment format="DD MMM YY">{post.created_at}</Moment> &middot; <Link to={"/posts/" + post.id}>{post.title}</Link>
        </li>
      )}
    </ul>
  </div>
);

const About = (props) => (
  <div>
    <h2>Hi, I'm Mike.</h2>
    <p>
      I like to write code for the web. I like to write programs in
      various languages that draw beautiful fractals. I like to explore
      math, physics, logic, electronics, and music. I write ops tasks
      away with code for <a href="https://newrelic.com">New Relic</a>.
      <EndMark/>
    </p>
  </div>
);

const Index = (props) => (
  <div>
    <About/>
    {props.posts.length > 0 ? (<RecentPosts posts={props.posts}/>) : null}
  </div>
);

const PostIndex = (props) => (
  <div>
    <h2>Posts</h2>
    {props.posts.length <= 0 ? (
      <p>Please write some posts.</p>
    ) : (
      <ul className="post-list">
        {props.posts.map((post) =>
          <li key={post.id.toString()}>
            <Moment format="DD MMM YY">{post.created_at}</Moment> &middot; <Link to={"/posts/" + post.id}>{post.title}</Link>
          </li>
        )}
      </ul>
    )}
  </div>
);

class Admin extends Component {
  constructor(props) {
    super(props);

    // TODO check for token in browser storage.
    // TODO validate & refresh token?
    // Does this need to be handled asynchronously?

    this.token_path = 'io.hix.auth.token';
    
    this.storage = localStorage;
    //this.storage = sessionStorage;

    var token = this.storage.getItem(this.token_path) || null;

    this.state = {
      email: "",
      password: "",
      token: token,
      flash: null,
      flash_id: null,
      flash_class: null,
    };

    this.handleEmailChange = this.handleEmailChange.bind(this);
    this.handlePasswordChange = this.handlePasswordChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleLogOut = this.handleLogOut.bind(this);
  }
  // NOTE needs separate component.
  flash(message, class_name, delay) {
    if(!class_name) { class_name = 'info'; }
    if(!delay) { delay = 10000; }

    if(this.state.flash_id) {
      this.setState((prevState, props) => {
        clearTimeout(prevState.flash_id);
        return { flash_id: null };
      });
    }

    const id = window.setTimeout(() => {
      this.setState({
        flash: null,
        flash_class: null,
        flash_id: null
      });
    }, delay);

    this.setState({
      flash: message,
      flash_class: class_name,
      flash_id: id
    });

  }
  // NOTE react-validators? react-forms?
  handleEmailChange(event) {
    this.setState({email: event.target.value});
  }
  handlePasswordChange(event) {
    this.setState({password: event.target.value});
  }
  // NOTE push this state and events up into App?
  handleSubmit(event) {
    const uri = 'http://localhost:3000/api/v1/auth';

    axios.post(uri, {
      email: this.state.email,
      password: this.state.password
    }).then((res) => {
      var token = res.headers.authorization.split(' ')[1];
      this.storage.setItem(this.token_path, token);
      this.setState({token: token});
      this.flash('You have logged in.', 'success');
    }).catch((error) => {
      this.flash(`Log in failed: ${error.response.status}`, 'warning');
    });

    event.preventDefault();
  }
  handleLogOut() {
    this.setState({token: null});
    this.storage.removeItem(this.token_path);
    this.flash('You have logged out.', 'success');
  }
  render() {
    if(this.state.token) {
      return (
        <div>
          <h2>Admin</h2>
          {this.state.flash ? (
            <p className={this.state.flash_class}>{this.state.flash}</p>
          ) : (null)}
          <button className="pure-button" onClick={this.handleLogOut}>log out</button>
        </div>
      );
    } else {
      // Login Form
      return (
        <div>
          <h2>Hi.</h2>
          {this.state.flash ? (
            <p className={this.state.flash_class}>{this.state.flash}</p>
          ) : (null)}
          <form onSubmit={this.handleSubmit} className="pure-form pure-form-aligned" >
            <fieldset>
              <div className="pure-control-group">
                <label htmlFor="email"><i className="fa fa-envelope">&nbsp;</i></label>
                <input id="email" type="email" value={this.state.email} onChange={this.handleEmailChange} required />
                <span className="pure-form-message-inline"></span>
              </div>
              <div className="pure-control-group">
                <label htmlFor="password"><i className="fa fa-star">&nbsp;</i></label>
                <input id="password" type="password" value={this.state.password} onChange={this.handlePasswordChange} required />
                <span className="pure-form-message-inline"></span>
              </div>
              <div className="pure-control-group">
                <label>&nbsp;</label>
                <input type="submit" value="log in" className="pure-button" />
                <span className="pure-form-message-inline"></span>
              </div>
            </fieldset>
          </form>
        </div>
      );
    }
  };
};

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      posts: [],
      photos: [],
      photo_groups: [],
    };
  }
  componentDidMount() {
    this.fetchPhotos();
    this.fetchPosts();
  }
  fetchPhotos() {
    var uri = 'http://localhost:3000/api/v1/photos';

    axios.get(uri).then((res) => {
      var photos = {};
      var groups = {};

      res.data.forEach((photo) => {
        var date = new Date(photo.created_at);
        var group = moment(date).format("MMM YYYY");

        if(!groups[group]) {
          groups[group] = [];
        }
        photos[photo.id] = photo;
        groups[group] = groups[group].concat(photo);
      });

      this.setState({
        photos: photos,
        photo_groups: groups,
      });
    });
  }
  findPhoto(props) {
    const photo = this.state.photos[props.match.params.id];

    if(!photo) { return <NotFound/>; }
    return <Photo {...photo}/>;
  }
  fetchPosts() {
    var uri = 'http://localhost:3000/api/v1/posts';

    axios.get(uri).then((res) => {
      this.setState({ posts: res.data });
    });
  }
  findPost(props) {
    const id = props.match.params.id; 
    const post = this.state.posts.find((post) => (post.id.toString() === id));

    if(!post) { return <NotFound/>; }
    return <Post {...post}/>;
  }
  render() {
    return <Router>
      <div className="pure-g main">
        <div className="pure-u-1 pure-u-md-1-6"></div>
        <section className="pure-u-1 pure-u-md-1-6 nav-box">
          <Link to="/" className="logo">
            <img src="/images/knot.jpg" alt="hix.io" />
            <h1>hix.io</h1>
          </Link>
          <p>&#x2234;</p>
          <ul>
            <li><Link to='/posts'>posts</Link></li>
            <li><Link to='/photos'>photos</Link></li>
            <li><a href="https://code.hix.io">code</a></li>
          </ul>
          <p>&#x2235;</p>
          <p>&nbsp;</p>
          <p>&#x2234;</p>
          <ul>
            <li><a href="https://github.com/musl"><i className="fa fa-github"></i>&nbsp;musl</a></li>
            <li><a href="https://twitter.com/mooselix"><i className="fa fa-twitter"></i>&nbsp;@mooselix</a></li>
            <li><a href="https://freenode.net/"><i className="fa fa-terminal"></i>&nbsp;musl</a></li>
          </ul>
          <p>&#x2235;</p>
          <p>&nbsp;</p>
          <p>&#x2234;</p>
          <ul>
            <li><Link to='/admin'><i className="fa fa-key"></i></Link></li>
          </ul>
          <p>&#x2235;</p>
        </section>
        <section className="pure-u-1 pure-u-md-2-3 content-box">
          <div className="column-pad">
            <Switch>
              <Route path='/photos/:id'>{(props) => this.findPhoto(props)}</Route>
              <Route path='/photos'><PhotoIndex photos={this.state.photos} photo_groups={this.state.photo_groups}/></Route>
              <Route path='/posts/:id'>{(props) => this.findPost(props)}</Route>
              <Route path='/posts'><PostIndex posts={this.state.posts}/></Route>
              <Route path='/admin'><Admin {...this.state}/></Route>
              <Route path='/'><Index posts={this.state.posts}/></Route>
              <Route component={NotFound}/>
            </Switch>
          </div>
        </section>
      </div>
    </Router>;
  }
}

ReactDOM.render(<App/>, document.getElementById('root'));

// I know not with what weapons World War III will be fought, but World War IV will be fought with sticks and stones. -- Albert Einstein


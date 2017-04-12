// I know not with what weapons World War III will be fought, but World War IV will be fought with sticks and stones. -- Albert Einstein

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
    <a href={props.original_url}>
      <img src={props.original_url} alt={props.alt} className="pure-img photo-original"/>
    </a>
  </div>
);

const Thumbnail = (props) => (
  <Link to={"/photos/" + props.id} className="photo-thumbnail">
    <img src={props.thumbnail_url} alt={props.alt} className="photo-thumbnail"/>
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
          <h2><Moment format="MMMM YYYY">{group}</Moment></h2>
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

class App extends Component {
  constructor(props) {
    super(props)
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
        var key = new Date(date.getFullYear(), date.getMonth());

        if(!groups[key]) {
          groups[key] = [];
        }
        photos[photo.id] = photo;
        groups[key] = groups[key].concat(photo);
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
        </section>
        <section className="pure-u-1 pure-u-md-2-3 content-box">
          <div className="column-pad">
            <Switch>
              <Route path='/photos/:id'>{(props) => this.findPhoto(props)}</Route>
              <Route path='/photos'><PhotoIndex photos={this.state.photos} photo_groups={this.state.photo_groups}/></Route>
              <Route path='/posts/:id'>{(props) => this.findPost(props)}</Route>
              <Route path='/posts'><PostIndex posts={this.state.posts}/></Route>
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


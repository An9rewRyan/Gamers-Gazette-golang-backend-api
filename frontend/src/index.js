import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom";
import './index.css';
import ArticlesApp from './ArticlesApp';
import * as serviceWorker from './serviceWorker';

const Routing = () => {
  return(
    <Router>
      <Routes>
        {/* <Route path="/articles/*"  exact component={ArticlesApp} /> */}
        <Route exact path="/"  element={ArticlesApp} />
        {/* <Route exact path="games/" element={<GamesApp />}/> */}
      </Routes>
    </Router>
  )
}

ReactDOM.render(
  <React.StrictMode>
    <Routing />
  </React.StrictMode>,
  document.getElementById('root')
);

serviceWorker.unregister();
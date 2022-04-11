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
        <Switch>
          <Route path="/articles/*"  exact component={ArticlesApp} />
        {/* <Route exact path="games/" element={<GamesApp />}/> */}
        </Switch>
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
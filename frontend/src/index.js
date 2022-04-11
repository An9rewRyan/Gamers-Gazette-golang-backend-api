import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom";
import './index.css';
import ArticlesApp from './ArticlesApp';
import reportWebVitals from './reportWebVitals';
import * as serviceWorker from './serviceWorker';

const Routing = () => {
  return(
    <Router>
      <Routes>
        <Route exact path="articles/"  element={<ArticlesApp />} />
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

reportWebVitals();
serviceWorker.unregister();
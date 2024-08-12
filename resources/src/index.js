import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

const page = window.__INITIAL_PAGE__ || 'home';

ReactDOM.render(<App page={page} />, document.getElementById('root'));

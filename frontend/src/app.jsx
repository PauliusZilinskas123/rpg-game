import React from 'react';
import Navigation from './components/Navigation';
import Header from './components/Header';
import 'normalize.css';
import 'styles/index.scss';

const App = () => (
  <div className='App'>
  	<Header/>
    <Navigation/>
    <div>
      <h1>It Works!</h1>
      <p>This React project just works including <span className="redBg">module</span> local styles.</p>
      <p>This React project just works including <span className="testBg">module</span> local styles.</p>
      <p>Enjoy!</p>
    </div>
  </div>
);

export default App;

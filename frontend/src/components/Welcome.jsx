import React from 'react';
import ImageLoader from 'react-load-image'; 
import 'styles/index.scss';

function Preloader(props) {
  return <img src="spinner.gif" />;
}

class Welcome extends React.Component {
  render() {
    return(
      <ImageLoader
      src="quest.png" className="main-img"
    >
      <img />
      <div>Error!</div>
      <Preloader />
    </ImageLoader>
    );
  }
}

export default Welcome;

import React from 'react';

class Container extends React.Component {
    render() {
        return (
            <div className="cnt">
            <h1>It Works!</h1>
            <p>This React project just works including <span className="redBg">module</span> local styles.</p>
            <p>This React project just works including <span className="testBg">module</span> local styles.</p>
            <p>Enjoy!</p>
            </div>
        );
    }
}

export default Container;
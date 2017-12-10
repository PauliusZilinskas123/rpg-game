import React from 'react';
import styles from '../../src/components/test.scss';
//import styles from 'styles/common/button.scss';

console.log(styles);

const Header = () => (
  <div className="header">
    Header part <span className={styles.testBg}>test</span>
  </div>
);

export default Header;

import React from 'react';
import { Layout } from 'antd';

const { Footer } = Layout;

const FooterComponent: React.FC = () => {
  return (
    <Footer style={{ textAlign: 'center' }}>
      Built on <a href="https://github.com/alphaokxyz/qaweb">Qaweb</a> - the open-source software that powers Q&A communities.
      <br />
      © 2023 Qaweb.
    </Footer>
  );
};

export default FooterComponent;

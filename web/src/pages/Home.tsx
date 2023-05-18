import React from 'react';
import { Layout, theme } from 'antd';
import Header from './Header';
import Footer from './Footer';
import QuestionList from './Homequestionlist'; // 引入 QuestionList 组件

const { Content } = Layout;

const App: React.FC = () => {
  const {
    token: { colorBgContainer },
  } = theme.useToken();

  return (
    <Layout className="layout">
      <Header />
      <Content>
        <div className="site-layout-content" style={{ background: colorBgContainer }}>
          <QuestionList /> {/* 替换为 QuestionList 组件 */}
        </div>
      </Content>
      <Footer />
    </Layout>
  );
};

export default App;

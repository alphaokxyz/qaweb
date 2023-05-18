import React from 'react';
import { Layout, Button } from 'antd';
import { Link } from 'react-router-dom';

const { Header } = Layout;

interface TokenData {
  username: string;
  exp: number;
  iss: string;
}

const HeaderComponent: React.FC = () => {
  const token = localStorage.getItem('token');
  let isLoggedIn = false;
  let username = '';

  if (token) {
    isLoggedIn = true;
    const tokenData = JSON.parse(atob(token.split('.')[1])) as TokenData;
    username = tokenData.username;
  }

  const handleLogout = () => {
    localStorage.removeItem('token');
  };

  return (
    <Header style={{ background: 'transparent', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
      <div style={{ display: 'flex', alignItems: 'center' }}>
        <Link to="/">
          <img src="/favicon.ico" alt="Logo" style={{ width: '80px', height: '80px', marginRight: '10px' }} />
        </Link>
      </div>
      <div>
        {isLoggedIn ? (
          <>
            <span style={{ color: '#000', marginRight: '8px' }}>{username}</span>
            <Button type="primary" onClick={handleLogout} style={{ marginRight: '8px' }}>
              注销
            </Button>
            <Button type="primary">
              <Link to="/new" style={{ color: '#fff' }}>新建</Link>
            </Button>
            <Button type="primary">
              <Link to="/users" style={{ color: '#fff' }}>用户列表</Link>
            </Button>
          </>
        ) : (
          <>
            <Button type="primary">
              <Link to="/login" style={{ color: '#fff' }}>登录</Link>
            </Button>
            <Button type="primary">
              <Link to="/signup" style={{ color: '#fff' }}>注册</Link>
            </Button>
          </>
        )}
      </div>
    </Header>
  );
};

export default HeaderComponent;

import  { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Form, Input, Button, message } from 'antd';
import axios from 'axios';
import { api } from '../config';

const LoginForm = () => {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (values: { username: string, password: string }) => {
    try {
      setLoading(true);

      const response = await axios.post(`${api}login`, {
        username: values.username,
        password: values.password,
      });

      const { token, status, message: responseMessage } = response.data;

      if (status === 200) {
        // 将令牌保存到本地存储或cookie中
        localStorage.setItem('token', token);

        message.success('登录成功');
        navigate('/'); // 导航到仪表盘页面或其他受保护的页面
      } else if (status === 1003) {
        message.error(responseMessage);
      } else if (status === 1002) {
        message.error(responseMessage);
      }
    } catch (error) {
      setLoading(false);
      message.error('登录失败，请检查用户名和密码');
    }
  };

  return (
    <Form onFinish={handleLogin}>
      <Form.Item
        name="username"
        rules={[{ required: true, message: '请输入用户名' }]}
      >
        <Input placeholder="用户名" />
      </Form.Item>

      <Form.Item
        name="password"
        rules={[{ required: true, message: '请输入密码' }]}
      >
        <Input.Password placeholder="密码" />
      </Form.Item>

      <Form.Item>
        <Button type="primary" htmlType="submit" loading={loading}>
          登录
        </Button>
      </Form.Item>
    </Form>
  );
};

export default LoginForm;

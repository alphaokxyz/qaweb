import React, { useEffect, useState } from 'react';
import { Table } from 'antd';
import axios from 'axios';
import { api } from '../config';

interface User {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  username: string;
  password: string;
  role: number;
}

interface UserData {
  data: User[];
  message: string;
  status: number;
  total: number;
}

const UserList: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [total, setTotal] = useState<number>(0);

  useEffect(() => {
    // 在组件挂载时获取用户数据
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    try {
      const response = await axios.get<UserData>(`${api}users`);
      const userData = response.data;
      setUsers(userData.data);
      setTotal(userData.total);
    } catch (error) {
      console.error('Failed to fetch users:', error);
    }
  };

  const columns = [
    {
      title: '用户名',
      dataIndex: 'username',
      key: 'username',
    },
    {
      title: '注册时间',
      dataIndex: 'CreatedAt',
      key: 'CreatedAt',
    },
  ];

  return (
    <div>
      <h2>用户数: {total}</h2>
      <Table dataSource={users} columns={columns} rowKey="ID" />
    </div>
  );
};

export default UserList;

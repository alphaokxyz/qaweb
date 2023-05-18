import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Input, Button, message } from 'antd';
import axios from 'axios';
import { useLocalObservable, observer } from 'mobx-react-lite';
import { api } from '../config';

type AddQuestionProps = {};

const AddQuestion: React.FC<AddQuestionProps> = observer(() => {
  const [title, setTitle] = useState<string>('');
  const [questionDetail, setQuestionDetail] = useState<string>('');
  const navigate = useNavigate();

  const store = useLocalObservable(() => ({
    get isLoggedIn() {
      return localStorage.getItem('token') !== null;
    },
  }));

  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };

  const handleQuestionDetailChange = (
    e: React.ChangeEvent<HTMLTextAreaElement>
  ) => {
    setQuestionDetail(e.target.value);
  };

  const handleAddQuestion = async () => {
    try {
      const response = await axios.post(
        `${api}question/add`,
        {
          title: title,
          questiondetail: questionDetail,
        },
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        }
      );

      if (response.status === 200) {
        message.success('问题添加成功');
        navigate('/');
      } else {
        message.error('问题添加失败');
      }
    } catch (error) {
      message.error('问题添加失败');
    }
  };

  if (!store.isLoggedIn) {
    return <div>请登录后提问</div>;
  }

  return (
    <div>
      <Input
        value={title}
        onChange={handleTitleChange}
        placeholder="请输入问题标题"
      />
      <Input.TextArea
        value={questionDetail}
        onChange={handleQuestionDetailChange}
        placeholder="请输入问题详情"
      />
      <Button type="primary" onClick={handleAddQuestion}>
        发送
      </Button>
    </div>
  );
});

export default AddQuestion;

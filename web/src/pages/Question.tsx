import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { observer } from 'mobx-react-lite';
import { Spin, Card, Button, Input, message, Layout } from 'antd';
import axios from 'axios';
import QuestionDetails from './Comment';
import { api } from '../config';
import Header from './Header';
import Footer from './Footer';

const { Content } = Layout;

interface QuestionData {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  title: string;
  questiondetail: string;
  username: string;
}

const QuestionDetail = observer(() => {
  const { id } = useParams();
  const [loading, setLoading] = useState(true);
  const [question, setQuestion] = useState<QuestionData | null>(null);
  const [comment, setComment] = useState('');

  useEffect(() => {
    const fetchData = async () => {
      try {
        if (!id) {
          return;
        }
        const response = await axios.get(`${api}question/info/${id}`);
        setQuestion(response.data.data);
        setLoading(false);
      } catch (error) {
        console.error(error);
        setLoading(false);
      }
    };

    fetchData();
  }, [id]);

  const handleCommentSubmit = async () => {
    const token = localStorage.getItem('token');
    if (!token) {
      message.error('请登录后再发表评论！');
      return;
    }

    const requestData = {
      aid: id ? parseInt(id) : 0,
      answerdetail: comment,
    };

    try {
      const response = await axios.post(`${api}answer/add`, requestData, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      message.success('评论已发送！');
      setComment('');

      // 使用 response 的值
      console.log(response.data);
    } catch (error) {
      console.error(error);
      message.error('评论发送失败，请重试！');
    }
  };

  if (loading) {
    return <Spin />;
  }

  if (!question) {
    return null; // 或者渲染一个加载失败的错误信息
  }

  // 格式化CreatedAt字段
  const formattedCreatedAt = new Date(question.CreatedAt).toLocaleString('en-US', {
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
    hour12: false,
    hour: 'numeric',
    minute: 'numeric',
    second: 'numeric',
  });

  return (
    <Layout>
      <Header />
      <Content>
        <Card title={question.title} style={{ width: 300 }}>
          <p>{question.username}</p>
          <p>{question.questiondetail}</p>
          <p>{formattedCreatedAt}</p> {/* 显示格式化后的CreatedAt字段 */}
        </Card>
        <QuestionDetails questionId={id || ''} />
        {localStorage.getItem('token') ? (
          <>
            <Input.TextArea
              rows={4}
              value={comment}
              onChange={(e) => setComment(e.target.value)}
            />
            <Button type="primary" onClick={handleCommentSubmit}>
              发送评论
            </Button>
          </>
        ) : (
          <p>请登录后再发表评论</p>
        )}
      </Content>
      <Footer />
    </Layout>
  );
});

export default QuestionDetail;

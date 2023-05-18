import React, { useEffect, useState } from 'react';
import { List, Pagination } from 'antd';
import axios from 'axios';
import { api } from '../config';

interface Question {
  ID: number;
  title: string;
  username: string;
  CreatedAt: string;
}

const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  const formattedDate = `${date.toLocaleDateString('en-US')} ${date.toLocaleTimeString('en-US', { hour12: false })}`;
  return formattedDate;
};

const QuestionList: React.FC = () => {
  const [questions, setQuestions] = useState<Question[]>([]);
  const [total, setTotal] = useState<number>(0);
  const [total2, setTotal2] = useState<number[]>([]);
  const [currentPage, setCurrentPage] = useState<number>(1);

  useEffect(() => {
    fetchQuestions();
  }, [currentPage]);

  const fetchQuestions = async () => {
    const pageSize = 5;
    try {
      const response = await axios.get(`${api}questions`, {
        params: {
          pagesize: pageSize,
          pagenum: currentPage,
        },
      });

      const { data } = response.data;
      setQuestions(data);
      setTotal(response.data.total);

      const answerRequests = data.map(async (question: Question) => {
        const response2 = await axios.get(`${api}answers/${question.ID}`);
        return response2.data.total;
      });
      const total2Values = await Promise.all(answerRequests);
      setTotal2(total2Values);
    } catch (error) {
      console.error('Error fetching questions:', error);
    }
  };

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  return (
    <div>
      <List
        itemLayout="vertical"
        dataSource={questions}
        renderItem={(question: Question, index: number) => (
          <List.Item key={question.ID}>
            <List.Item.Meta
              title={<a href={`/question/${question.ID}`}>{question.title}</a>}
              description={`${question.username} ${total2[index]}`}
            />
            <div>{formatDate(question.CreatedAt)}</div>
          </List.Item>
        )}
      />
      <Pagination
        current={currentPage}
        pageSize={5}
        total={total}
        onChange={handlePageChange}
        showSizeChanger={false}
      />
    </div>
  );
};

export default QuestionList;

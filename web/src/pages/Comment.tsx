import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { api } from '../config';

interface QuestionDetailsProps {
  questionId: string;
}

interface Question {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: null;
  title: string;
  questiondetail: string;
  username: string;
}

interface Answer {
  Question: Question;
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: null;
  aid: number;
  answerdetail: string;
  username: string;
}

const QuestionDetails: React.FC<QuestionDetailsProps> = ({ questionId }) => {
  const [question, setQuestion] = useState<Question | null>(null);
  const [answers, setAnswers] = useState<Answer[]>([]);

  useEffect(() => {
    const fetchQuestionAndAnswers = async () => {
      try {
        const response = await axios.get(`${api}answers/${questionId}`);
        const { data } = response;
        if (data.data.length > 0) {
          setQuestion(data.data[0].Question);
          setAnswers(data.data);
        } else {
          setQuestion(null);
          setAnswers([]);
        }
      } catch (error) {
        console.error('Error fetching question and answers:', error);
      }
    };

    fetchQuestionAndAnswers();
  }, [questionId]);

  if (!question) {
    return <div>暂无评论</div>;
  }

  // 格式化时间函数，改为美国24小时制
  const formatTime = (timeString: string) => {
    const date = new Date(timeString);
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    const hour = date.getHours() < 10 ? `0${date.getHours()}` : date.getHours();
    const minute = date.getMinutes() < 10 ? `0${date.getMinutes()}` : date.getMinutes();
    const second = date.getSeconds() < 10 ? `0${date.getSeconds()}` : date.getSeconds();
    return `${month}/${day}/${year} ${hour}:${minute}:${second}`;
  };

  return (
    <div>
      <h3>{question.title}</h3>
      {answers.map((answer) => (
        <div key={answer.ID}>
          <p>{answer.username}</p>
          <p>{answer.answerdetail}</p>
          <p>{formatTime(answer.CreatedAt)}</p> {/* 使用格式化时间函数 */}
          <hr />
        </div>
      ))}
    </div>
  );
};

export default QuestionDetails;

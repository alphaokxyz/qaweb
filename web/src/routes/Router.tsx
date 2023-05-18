import { createBrowserRouter } from 'react-router-dom';
import Home from '../pages/Home';
import Login from '../pages/Login';
import Register from '../pages/Register';
import Notfound from '../pages/Notfound';
import QuestionDetail from '../pages/Question';
import NewQuestion from '../pages/NewQuestion';
import Users from '../pages/Users';

const Router = createBrowserRouter([
  { path: '/', element: <Home /> },
  { path: '/login', element: <Login /> },
  { path: '/signup', element: <Register /> },
  { path: '/new', element: <NewQuestion /> },
  { path: '/users', element: <Users /> },
  
  { path: '/question/:id', element: <QuestionDetail /> }, // 添加动态路由参数 ":id"
  { path: '/404', element: <Notfound /> },
  { path: '*', element: <Notfound /> },
]);

export default Router;

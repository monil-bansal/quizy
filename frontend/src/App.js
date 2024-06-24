import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import './App.css';
import Quiz from './components/Quiz';
import QuizList from './components/QuizList';
import NotExist from './components/NotExist';

function App() {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={QuizList} />
        <Route path="/quiz/*" component={Quiz} />
        <Route path="*" component={NotExist} />
      </Switch>
    </Router>
  );
}

export default App;

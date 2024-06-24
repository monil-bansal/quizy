import { useState, useEffect } from 'react';
import QuizTile from './QuizTile';

function QuizList() {
    const [quizList, setQuizList] = useState([]);
    useEffect(() => {
        fetch('https://localhost:443/quiz', {
            method: 'GET',
          })
           .then((response) => {
                return response.json()
            })
           .then((data) => {
              setQuizList([])
              for(let i = 0; i<data.length;i++){
                setQuizList(oldArray => [...oldArray, <QuizTile key={i} quizTitle={data[i].title} quizId={data[i].quizId}/>]);
              }
           })
           .catch((err) => {
              console.log(err.message);
           });
     }, []);
  return (
    <div className="Quiz">
        <h1> All Quizes present </h1>
       {quizList}
    </div>
  );
}

export default QuizList;

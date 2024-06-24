import Question from "./Question";
import { useState, useEffect } from 'react';

function Quiz() {
  const [quiz, setQuiz] = useState(null);
  useEffect(() => {

      let pathSplit = window.location.href.split("/")
      let quizId = pathSplit[pathSplit.length - 1];
    
      fetch('https://localhost:443/quiz/' + quizId, {
          method: 'GET',
        })
         .then((response) => {
              return response.json()
          })
         .then((data) => {
          console.log(data)
          setQuiz(data)
         })
         .catch((err) => {
            console.log(err.message);
         });
   }, []);

  const rows = [];
  for (let i = 0; i < 5; i++) {
      rows.push(<Question questionId = {i}/>);
  }

  return (
    <div className="Quiz">
      <h3> {quiz.title} </h3>
       {rows}
       <button>submit quiz</button>
    </div>
  );
}

export default Quiz;

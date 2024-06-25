import Question from "./Question";
import { useState, useEffect } from 'react';
import {FidgetSpinner} from 'react-loader-spinner'

function Quiz() {
  const [quiz, setQuiz] = useState(null);
  const [isLoading, setLoading] = useState(true);
  const [questions, setQuestions] = useState([]);

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

        for (let i = 0; i < data.questions.length; i++) {
          setQuestions(oldArray => [...oldArray, <Question key={i} questionId={i} options={data.questions[i].options}  question={data.questions[i].question}/>]);

        }
        setLoading(false)
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  // for (let i = 0; i < 5; i++) {
  //     rows.push(<Question questionId = {i}/>);
  // }
  if (isLoading) {
    return (
      <div className="loadingContainer">
        <FidgetSpinner
          visible={true}
          height="80"
          width="80"
          ariaLabel="fidget-spinner-loading"
          wrapperStyle={{}}
          wrapperClass="fidget-spinner-wrapper"
        />
      </div>
    )
  } else {

    return (
      <div className="Quiz">
        <h3> {quiz.title} </h3>
        {questions}
        <button>submit quiz</button>
      </div>)
  }


  // return (

  // );
}

export default Quiz;

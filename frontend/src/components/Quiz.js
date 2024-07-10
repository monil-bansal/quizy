import Question from "./Question";
import { useState, useEffect } from 'react';
import {FidgetSpinner} from 'react-loader-spinner'
import NotExist from "./NotExist";

function loadingComponent(){
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
}

function submitQuiz(quizId, questions, setScore){
  var l =  questions.length
  var answers = []
  for(let i = 0;i<l;i++){
    var ele = document.getElementsByName(i);
    var answered = false;
    for (let j = 0; j < ele.length; j++) {
      if (ele[j].checked){
        answered = true;
        answers.push({"answer": j});
      }
    }
    if(!answered){
      // Thought: can use in future to have different scoring for unanswered and wrong answered => like JEE MAINS.
      answers.push({"answer" : -1});
    }
  }
  fetch('https://localhost:443/submit/' + quizId, {
    method: 'POST',
    body: JSON.stringify({
      "questions": answers
    })
  }).then((response) => {
    return response.json()
  })
  .then((data) => {
    setScore(data)
    return data
  })
  .catch((err) => {
    console.log(err.message);
  });
}

function quizComponent(quizId, title, questions, setScore) {
  return (
    <div className="Quiz">
      <h3> {title} </h3>
      {questions}
      <button onClick={() => submitQuiz(quizId, questions, setScore)}>submit quiz</button>
    </div>)
}

function Quiz() {
  const [quiz, setQuiz] = useState(null);
  const [notExists, setNotExists] = useState(false);
  const [isLoading, setLoading] = useState(true);
  const [questions, setQuestions] = useState([]);
  const [score, setScore] = useState(-1)

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
        setQuiz(data)
        if(data.questions == null){
          setQuestions([])
          setNotExists(true)
          return
        }
        for (let i = 0; i < data.questions.length; i++) {
          setQuestions(oldArray => [...oldArray, <Question key={i} questionId={i} options={data.questions[i].options}  question={data.questions[i].question}/>]);

        }
        setLoading(false)
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  if (notExists) {
    return (<NotExist />)
  }
  else if (isLoading) {
    return loadingComponent()
  } 
  else if(score!= -1){
    return (<p> Scoore : {score} </p>)
  }
  else {
    let pathSplit = window.location.href.split("/")
    let quizId = pathSplit[pathSplit.length - 1];
    return quizComponent(quizId, quiz.title, questions, setScore)
  }
}

export default Quiz;

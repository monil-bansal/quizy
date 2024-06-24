import { Redirect } from 'react-router-dom';
import { useHistory } from "react-router-dom";

function QuizTile(props) {
    let navigate = useHistory(); 
    const routeChange = () =>{ 
      let path = `/quiz/` + props.quizId; 
      navigate.push(path);
    }
    return (
      <div className="quizTile">
        <fieldset>
          <h4> {props.quizTitle} </h4>
          <p> quiz description </p>
          <button onClick={routeChange}>start quiz</button>
        </fieldset>
      </div>
    );
  }
  
export default QuizTile;
  
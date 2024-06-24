function Question(props) {
    var questionId = props.questionId
    return (
      <div className="Question">
        <fieldset>
          <legend>Question:</legend>
          
          <div>
            <input type="radio" id={"huey" + questionId} name={questionId} value={"huey" + questionId} />
            <label for={"huey" + questionId}>Huey</label>
          </div>

          <div>
            <input type="radio" id={"dewey" + questionId} name={questionId} value={"dewey" + questionId} />
            <label for={"dewey" + questionId}>Dewey</label>
          </div>

          <div>
            <input type="radio" id={"louie" + questionId} name={questionId} value={"louie" + questionId} />
            <label for={"louie" + questionId}>Louie</label>
          </div>
        </fieldset>
      </div>
    );
  }
  
  export default Question;
  
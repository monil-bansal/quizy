function Question(props) {
  var questionId = props.questionId
  let options = [];
  for (let i = 0; i < props.options.length; i++) {
    let id = questionId*100 + i; // generating unique id in a quiz
    options.push(<div key={i}>
      <input type="radio" id={id} name={questionId} value={id} />
      <label htmlFor={id}>{props.options[i]}</label>
    </div>)
  }
  return (
    <div className="Question">
      <fieldset>
        <legend>{props.question}</legend>
        {options}
      </fieldset>
    </div>
  );
}

export default Question;

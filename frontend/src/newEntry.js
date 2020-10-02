import React from 'react';
import './App.css';
import ApiService from './service'

function NewEntry(props) {
  let state = props.state;
  let setState = props.setState;

  const clearState = () => {
    setState({ systolic: "", diastolic: "", heartrate: "" })
  }

  const setSystolic = (evt) => {
    setState({ ...state, systolic: evt.target.value })
  }

  const setDiastolic = (evt) => {
    setState({ ...state, diastolic: evt.target.value })
  }

  const setHeartrate = (evt) => {
    setState({ ...state, heartrate: evt.target.value })
  }

  const handleSubmit = (evt) => {
    evt.preventDefault();
    if(ApiService.postNewEntry(state)) {
      clearState()
    } else {
      console.log("FAIL")
    }
  }

  return(
    <div>
      <form id="newEntryForm" onSubmit={handleSubmit}>
        <input type="text" value={state.systolic} name="systolic" placeholder="sys" onChange={setSystolic} />
        <input type="text" value={state.diastolic} name="diastolic" placeholder="dia" onChange={setDiastolic} />
        <input type="text" value={state.heartrate} name="heartrate" placeholder="hr" onChange={setHeartrate} />
        <input type="submit" value="Submit" />
      </form>
    </div>
  )
}

export default NewEntry;

import axios from 'axios';

const domain = 'http://localhost:9000'

const ApiService = {
  async postNewEntry(state) {
    console.log(state)
    axios.post(domain + '/entries', state)
    .then(function (response) {
      console.log(response)
      return true;
    })
    .catch(function (error) {
      console.log(error)
      return false;
    })
  },

  async getAllEntries() {
    console.log("Getting all entries...")
    axios.get(domain + '/entries')
    .then(function (response) {
      console.log(response)
    })
    .catch(function (error) {
      console.log(error)
    })
  }
}

export default ApiService;

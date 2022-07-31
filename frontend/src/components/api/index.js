import axios from 'axios';

const axiosOptions = {
  headers: {
  'Content-Type': 'application/json'
  },
  mode: 'same-origin', // no-cors, *cors, same-origin
  cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
  credentials: 'same-origin', // include, *same-origin, omit
}

const url = '/diagnosis/v1/index.html';

export const fetchPosts = (newPost) => axios.post(url, newPost, axiosOptions);

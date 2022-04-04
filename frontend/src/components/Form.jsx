import { useEffect } from "react"
import axios from 'axios'


const getData = async () => {

  // const params = new URLSearchParams();
  // params.append('gen', 'female');
  // params.append('age', 28);
  // params.append('symps', 'Android Developer');

  // const config = {
  //   headers: {
  //     'Content-Type': 'application/json'
  //   }
  // }

  // const json = JSON.stringify({ gen: 'female', age: 28, symps : 'Android Developer' });

    
  // const res = await axios.post("http://localhost:8080/diagnosis/v1/index.html", json);
  // console.log(res)
  // const data = await res.json();
  // console.log(data)
  // console.log(res)
    // console.log(body)
    const api = await axios('http://localhost:8080/diagnosis/v1/index.html',{
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'same-origin', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
        // 'Content-Type': 'application/json'
        'Content-Type': 'application/json',
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify({ gen: 'female', age: 28, symps : 'Android Developer' }) // body data type must match "Content-Type" header
        });

    // const data = await api.json();
    // console.log(data)
};

function Form() {

    useEffect(() => {
        getData();
    }, [])
    
  return (
    <div>Form</div>
  )
}

export default Form
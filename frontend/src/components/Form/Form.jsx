import { Button, FormControl, FormControlLabel, FormLabel, Grid, Paper, Radio, RadioGroup, TextField } from '@material-ui/core'
import { useEffect, useState } from 'react'
import useStyles from './styles'
import Tags from './Tags';
import axios from 'axios';

const defaultValues = {
  age: 0,
  gen: "female",
  symps: []
};

const GenderComponent = ({formValues, handleInputChange}) => {
  return (
    <FormControl>
      <FormLabel>Gender</FormLabel>
      <RadioGroup
        name="gen"
        value={formValues.gen || ''}
        onChange={handleInputChange}
        row
      >
        <FormControlLabel
          key="male"
          value="male"
          control={<Radio size="small" />}
          label="Male"
        />
        <FormControlLabel
          key="female"
          value="female"
          control={<Radio size="small" />}
          label="Female"
        />
      </RadioGroup>
    </FormControl>
  )
}

const AgeComponent = ({formValues, handleInputChange}) => {
  return (
    <TextField
      id="age-input"
      name="age"
      label="Age"
      type="number"
      value={formValues.age}
      onChange={handleInputChange}
      min={0}
    />
  )
}

function Form() {
    
  const classes = useStyles()
  
  const [formValues, setFormValues] = useState(defaultValues);
  
  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormValues({
      ...formValues,
      [name]: value,
    });
  };

  const [resData, setResData] = useState({
    diseases: [
      {
        name: "",
        symptoms:""
      }
  ]})

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post(`http://${process.env.DIAGAPIIP}:${process.env.DIAGAPIPORT}/diagnosis/v1/index.html`, {...formValues, age: parseInt(formValues.age)}, {
      headers: {
      'Content-Type': 'application/json'
      },
      mode: 'same-origin', // no-cors, *cors, same-origin
      cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
      credentials: 'same-origin', // include, *same-origin, omit
    })
    .then(res => setResData({...resData, ...res.data}))
    
  };

  useEffect(() => {
    console.log(resData)
  }, [resData])
  

  return (
    <Paper className={classes.paper}>
      <form autoComplete='off' noValidate className={classes.form} onSubmit={handleSubmit}>
        <Grid container alignItems="center" justifyContent="center" direction="column">
          <Grid item xs={3}>
            <AgeComponent formValues={formValues} handleInputChange={handleInputChange} style={{margin: "20px"}}/>
          </Grid>
          <Grid itemxs={3}>
            <GenderComponent formValues={formValues} handleInputChange={handleInputChange} style={{padding: '5px'}}/>
          </Grid>
          <Grid itemxs={3}>
            <Tags formValues={formValues} setFormValues={setFormValues}/>
          </Grid>
          <Button variant="contained" color="primary" type="submit" xs={3}>
            Submit
          </Button>
        </Grid>
      </form>

      <Grid container alignItems="center" justifyContent="center" direction="column">
          <div>{undefined || resData["diseases"][0].name}</div>
          {/* <div>{undefined || resData["diseases"][1].name}</div> */}

          {resData.diseases.map((item, index) => {
            <Grid item key={index}>{item.name}</Grid>
          })}
      </Grid>
    </Paper>
  )
}

export default Form
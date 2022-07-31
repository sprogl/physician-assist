import React, { useState } from 'react'
import useStyles from './styles'
import Tags from './Tags'
import CheckboxesTags from './TagsAutoComplete'
import { Accordion, AccordionDetails, AccordionSummary, Button, FormControl, FormControlLabel, FormLabel, Grid, List, ListItem, ListItemIcon, ListItemText, Paper, Radio, RadioGroup, TextField, Typography } from '@material-ui/core'
import ExpandMoreIcon from '@material-ui/icons/ExpandMore'
import CheckCircleOutlineIcon from '@material-ui/icons/CheckCircleOutline'
import { fetchPosts } from '../api'

const defaultValues = {
  age: 11,
  gen: "female",
  symps: ['Goh Gije']
};

const SymptomList = React.memo(function SymptomList(props) {
  const {sympName} = props
  return (
    <List component="nav" aria-label="contacts">
      <ListItem button>
        <ListItemIcon>
          <CheckCircleOutlineIcon />
        </ListItemIcon>
        <ListItemText primary={sympName} />
      </ListItem>
    </List>
  );
})

const AccordionSection = (props) => {
  const [ isExpanded, setIsExpanded ] = useState(true)
  const { item } = props

  const handleAccordionExpanded = () => {
    setIsExpanded(!isExpanded)
  }
  return (
    <Accordion onClick={handleAccordionExpanded} expanded={isExpanded}>
      <AccordionSummary
        expandIcon={<ExpandMoreIcon />}
        aria-controls="panel1a-content"
        id="panel1a-header"
      >
        <Typography sx={{fondSize: '15rem'}}>{item.name}</Typography>
      </AccordionSummary>
      <AccordionDetails>
        <div>
          {item.symptoms && item.symptoms.map((symp, i) => <SymptomList key={i} sympName={symp} /> )}
        </div>
      </AccordionDetails>
    </Accordion>
  )}

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
  
  const [formValues, setFormValues] = useState(defaultValues)
  const [isNewData, setIsNewData] = useState(false)
  const handleInputChange = (e) => {
    const { name, value } = e.target
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

  const handleSubmit = async e => {
    e.preventDefault();
    const data = await fetchPosts({...formValues, age: parseInt(formValues.age)})
    setResData({...resData, ...data.data})
    setIsNewData(true)  
  };
  

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
          {/* <Grid itemxs={3}>
            <CheckboxesTags />
          </Grid> */}
          <Button variant="contained" color="primary" type="submit" xs={3}>
            Submit
          </Button>
        </Grid>
      </form>
      {isNewData && resData["diseases"].map((item, index) => 
        <AccordionSection key={index} item={item} />
      )}
    </Paper>
  )
}

export default Form
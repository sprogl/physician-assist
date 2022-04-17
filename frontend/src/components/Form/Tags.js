import React, { useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Chip from '@material-ui/core/Chip';
import { ListSubheader } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
    maxWidth: 400,
  },
  chips: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  chip: {
    margin: 2,
  },
}));

export default function MultipleSelect({formValues, setFormValues}) {
  const classes = useStyles();

  const [personName, setPersonName] = useState([]);
  const handleChange = (event) => {
    setPersonName(event.target.value);
    setFormValues({
      ...formValues,
      [event.target.name]: event.target.value,
    })
  };

  const renderSelectGroup = (cats, index) => {
    const items = cats.symps.map((p) => {
      return (
        <MenuItem key={p} value={p}>
          {p}
        </MenuItem>
      );
    });
    return [<ListSubheader key={index}>{cats.category}</ListSubheader>, items];
  };
  
  return (
    <div>
      <FormControl className={classes.formControl}>
        <InputLabel id="select-multiple-chip">Symptoms</InputLabel>
          <Select
            labelId="select-multiple-chip"
            id="select-multiple-chip"
            multiple
            name='symps'
            value={personName}
            onChange={handleChange}
            input={<Input id="select-multiple-chip" />}
            renderValue={(selected) => (
              <div className={classes.chips}>
                {selected.map((value) => (
                  <Chip key={value} label={value} className={classes.chip} />
                ))}
              </div>
            )}
          >
            {Symptoms?.map((p, index) => renderSelectGroup(p, index))}
          </Select>
      </FormControl>
    </div>
  );
}

const Symptoms = [
  {
    category: 'My ... hearts:',
    symps: [
      'Stomache', 'Head', 'Abdomen', 'Back', 'Chest',
      'Ear', 'Pelvis', 'Tooth', 'Rectum', 'Skin',
      'Leg', 'Chronic pain'
    ]
  },
  {
    category: 'I feel:',
    symps: [
      'Chills', 'Fever', 'Paresthesia (numbness, tingling, electric tweaks)',
      'Light-headed', 'Dizzy', 'Dizzy – about to black out', 'Dizzy – with the room spinning around me', 
      'My mouth is dry', 'Nauseated', 'Sick', 'like I have the flu', 'like I have to vomit', 'Short of breath',
      'Sleepy', 'Sweaty', 'Thirsty', 'Tired', 'Weak'
    ]
  }
];
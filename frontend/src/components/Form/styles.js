import { makeStyles } from '@material-ui/core/styles';

export default makeStyles((theme) => ({
  root: {
    '& .MuiTextField-root': {
      margin: theme.spacing(1),
    },
  },
  paper: {
    padding: theme.spacing(2),
  },
  form: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  fileInput: {
    width: '97%',
    margin: '10px 0',
  },
  buttonSubmit: {
    marginBottom: 10,
  },
  TagsInput: {
    margin: theme.spacing(1),
    minWidth: 120,
    display: 'flex',
    flexWrap: 'wrap',
    maxWidth: 300
  },
  gender: {
    padding: theme.spacing(2),
    margin: theme.spacing(1),
  },
  chips: {
    display: 'flex',
    flexWrap: 'wrap',
    maxWidth: '300px'
  },
  chip: {
    margin: 2,
    maxWidth: 300,
    display: 'flex',
    flexWrap: 'wrap',
  },
}));
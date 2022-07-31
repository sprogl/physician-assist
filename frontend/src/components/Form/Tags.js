import React, { useState } from "react"
import { makeStyles } from "@material-ui/core/styles"
import Input from "@material-ui/core/Input"
import InputLabel from "@material-ui/core/InputLabel"
import MenuItem from "@material-ui/core/MenuItem"
import FormControl from "@material-ui/core/FormControl"
import Select from "@material-ui/core/Select"
import Chip from "@material-ui/core/Chip"
import Symptoms from './symptomList.json'

const useStyles = makeStyles((theme) => ({
	formControl: {
		margin: theme.spacing(1),
		minWidth: 120,
		maxWidth: 400,
	},
	chips: {
		display: "flex",
		flexWrap: "wrap",
	},
	chip: {
		margin: 2,
	},
}));

export default function MultipleSelect({ formValues, setFormValues }) {
	const classes = useStyles();

	const [symptomName, setSymptomName] = useState([Symptoms[Symptoms.length - 1]]);
	// const [symptomName, setSymptomName] = useState([formValues.symps]); // TODO: initialize the symptom from the Form component & not here

	const handleChange = (event) => {
		setSymptomName(event.target.value);
		setFormValues({
			...formValues,
			[event.target.name]: event.target.value,
		});
	};

	const renderSelectGroup = (cats) => {
		const items = cats.map((p) => {
			return (
				<MenuItem key={p} value={p}>
					{p}
				</MenuItem>
			);
		});
		return [items];
	};

	return (
		<div>
			<FormControl className={classes.formControl}>
				<InputLabel id="select-multiple-chip">Symptoms</InputLabel>
				<Select
					labelId="select-multiple-chip"
					id="select-multiple-chip"
					multiple
					name="symps"
					value={symptomName}
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
					{renderSelectGroup(Symptoms)}
				</Select>
			</FormControl>
		</div>
	);
}

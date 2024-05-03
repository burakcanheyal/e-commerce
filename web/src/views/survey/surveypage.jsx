import React, { useState } from 'react';
import { Typography, Button, RadioGroup, Radio, FormControlLabel, FormGroup, FormControl, FormLabel, Grid } from '@mui/material';

const Surveypage = () => {
  const [answers, setAnswers] = useState(Array(10).fill(''));

  const handleAnswerChange = (index, answer) => {
    const newAnswers = [...answers];
    newAnswers[index] = answer;
    setAnswers(newAnswers);
  };

  // questions
  const surveyQuestions = [
    "What is your favorite travel destination?",
    "What type of accommodation do you prefer when traveling?",
    "What is your preferred mode of transportation during travel?",
    "What is your favorite outdoor activity while traveling?",
    "What is the most important factor for you when choosing a travel destination?",
    "Do you prefer solo travel or traveling with companions?",
    "What is your favorite cuisine to try when traveling?",
    "What is the most memorable travel experience you've had?",
    "How do you plan your travel itinerary?",
    "What is your dream travel destination?",
  ];

  const renderSurveyQuestions = () => {
    return surveyQuestions.map((question, index) => (
        <Grid item key={index} xs={12}>
          <FormControl component="fieldset">
            <FormLabel component="legend">{`${index + 1}. ${question}`}</FormLabel>
            <RadioGroup
                value={answers[index]}
                onChange={(e) => handleAnswerChange(index, e.target.value)}
            >
              <FormControlLabel value="yes" control={<Radio />} label="Yes" />
              <FormControlLabel value="no" control={<Radio />} label="No" />
              <FormControlLabel value="maybe" control={<Radio />} label="Maybe" />
            </RadioGroup>
          </FormControl>
        </Grid>
    ));
  };

  const submitSurvey = () => {
    console.log(answers);
  };

  return (
      <div>
        <Typography variant="h4" gutterBottom>
          Travel Survey
        </Typography>
        <FormGroup>
          <Grid container spacing={2}>
            {renderSurveyQuestions()}
          </Grid>
        </FormGroup>
        <div>
        </div>
        <Button variant="contained" onClick={submitSurvey}>Submit</Button>
      </div>
  );
};

export default Surveypage;

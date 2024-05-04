import React, { useState, useEffect } from 'react';
import { Typography, Button, RadioGroup, Radio, FormControlLabel, FormGroup, Grid } from '@mui/material';
import axios from 'axios';

const Surveypage = () => {
  const [answers, setAnswers] = useState({});
  const [surveyQuestions, setSurveyQuestions] = useState([]);

  useEffect(() => {
    const fetchSurveyQuestions = async () => {
      try {
        const accessToken = localStorage.getItem('AccessToken');
        const response = await axios.get('http://localhost:8001/trips/', {
          headers: {
            'Authentication': `${accessToken}`
          }
        });
        setSurveyQuestions(response.data);
        // Initialize answers with 'yes' for each question
        const initialAnswers = {};
        response.data.forEach(question => {
          initialAnswers[question.description] = 'yes';
        });
        setAnswers(initialAnswers);
      } catch (error) {
        console.error('Error fetching survey questions:', error);
      }
    };

    fetchSurveyQuestions();
  }, []);

  const handleAnswerChange = (description, value) => {
    setAnswers({ ...answers, [description]: value });
  };

  const renderSurveyQuestions = () => {
    return surveyQuestions.map((question, index) => (
      <Grid item key={question.id} xs={12}>
        <Typography variant="body1" gutterBottom>
          {`${question.id}. ${question.description}`}
        </Typography>
        <RadioGroup
          value={answers[question.description]}
          onChange={(e) => handleAnswerChange(question.description, e.target.value)}
        >
          <FormControlLabel value="yes" control={<Radio />} label="Yes" />
          <FormControlLabel value="no" control={<Radio />} label="No" />
          <FormControlLabel value="maybe" control={<Radio />} label="Maybe" />
        </RadioGroup>
      </Grid>
    ));
  };

  const submitSurvey = async () => {
    try {
      const accessToken = localStorage.getItem('AccessToken');
      const postData = surveyQuestions.map((question, index) => ({
        id: 0,
        question_id: index + 1,
        points: answers[question.description] === 'yes' ? 2 : answers[question.description] === 'maybe' ? 1 : 0
      }));
      await axios.post('http://localhost:8001/trips/', postData, {
        headers: {
          'Authentication': `${accessToken}`
        }
      });
      console.log('Survey submitted successfully!');
    } catch (error) {
      console.error('Error submitting survey:', error);
    }
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
      <Button variant="contained" onClick={submitSurvey}>Submit</Button>
    </div>
  );
};

export default Surveypage;

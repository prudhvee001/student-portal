const express = require('express');
const app = express();

app.get('/attendance/:user_id', (req, res) => {
  const userId = req.params.user_id;
  const attendance = { user_id: userId, attendance_percentage: 90 };
  res.json(attendance);
});

app.listen(5003, () => {
  console.log('Attendance service running on port 5003');
});

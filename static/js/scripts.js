let startTime;

function startTracking() {
  const initialKills = document.getElementById("initialKills").value;
  if (initialKills !== "") {
    startTime = new Date();
    document.getElementById("killForm").style.display = "none";
    document.getElementById("stopForm").style.display = "block";
  }
}

function stopTracking() {
  const finalKills = document.getElementById("finalKills").value;
  if (finalKills !== "") {
    const endTime = new Date();
    const initialKills = document.getElementById("initialKills").value;
    const killsAchieved = finalKills - initialKills;
    const timeTaken = (endTime - startTime) / 1000; // time in seconds
    const averageTimePerKill = timeTaken / killsAchieved;

    const url = `/results?killsAchieved=${killsAchieved}&timeTaken=${timeTaken}&averageTimePerKill=${averageTimePerKill}`;
    window.location.href = url;
  }
}

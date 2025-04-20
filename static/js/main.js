let imgArray = [
  "static/img/rock.png",
  "static/img/paper.png",
  "static/img/scissors.png",
];

let selected = 0;

function choice(x) {
  selected = x;

  playerImg = document.getElementById("player_choice_image");

  if (selected == 0) {
    playerImg.src = imgArray[selected];
  } else if (selected == 1) {
    playerImg.src = imgArray[selected];
  } else {
    playerImg.src = imgArray[selected];
  }

  console.log(selected);
}

function play() {
  let checkbox = document.getElementById("flip-cpu");
  checkbox.checked = false;

  fetch("/play?c=" + selected)
    .then((response) => response.json())
    .then((data) => {
      if (selected == 0) {
        document.getElementById("player_choice").innerText =
          "El jugador eligio PIEDRA";
      } else if (selected == 1) {
        document.getElementById("player_choice").innerText =
          "El jugador eligio PAPEL";
      } else {
        document.getElementById("player_choice").innerText =
          "El jugador eligio TIJERA";
      }

      document.getElementById("player_score").innerText = data.player_score;
      document.getElementById("computer_score").innerText = data.computer_score;

      document.getElementById("computer_choice").innerText =
        data.computer_choice;
      document.getElementById("round_result").innerText = data.round_result;
      document.getElementById("round_message").innerText = data.message;

      var imgElement = document.getElementById("img_computer");
      imgElement.src = imgArray[data.computer_choice_int];
      checkbox.checked = true;
    });
}

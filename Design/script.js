const createPostBtn = document.querySelector(".createPostBtn");
const postModal = document.getElementById("postModal");
const confirmModal = document.getElementById("confirmModal");
const closePostModal = document.querySelector("#postModal .close");
const closeConfirmModal = document.querySelector("#confirmModal .close");
const publishBtn = document.getElementById("publishBtn");
const deleteBtn = document.getElementById("deleteBtn");
const continueBtn = document.getElementById("continueBtn");

createPostBtn.addEventListener("click", () => {
    postModal.style.display = "block";
});

closePostModal.addEventListener("click", () => {
    postModal.style.display = "none";
});

window.addEventListener("click", (event) => {
    if (event.target == postModal) {
        postModal.style.display = "none";
    }
});

publishBtn.addEventListener("click", () => {
    postModal.style.display = "none";
});

window.addEventListener("beforeunload", (event) => {
    if (postModal.style.display === "block") {
        event.preventDefault();
        confirmModal.style.display = "block";
    }
});

closeConfirmModal.addEventListener("click", () => {
    confirmModal.style.display = "none";
});

deleteBtn.addEventListener("click", () => {
    confirmModal.style.display = "none";
    postModal.style.display = "none";
});

continueBtn.addEventListener("click", () => {
    confirmModal.style.display = "none";
});

function previewProfilePic(event) {
    const reader = new FileReader();
    reader.onload = function(){
        const output = document.getElementById('profile-img-preview');
        output.src = reader.result;
    }
    reader.readAsDataURL(event.target.files[0]);
}

function previewProfilePic(event) {
    const reader = new FileReader();
    reader.onload = function() {
        const output = document.getElementById('profile-img-preview');
        output.src = reader.result;
    }
    reader.readAsDataURL(event.target.files[0]);
}


function addInterest() {
    const interestsContainer = document.getElementById('interests-inputs');
    const newInput = document.createElement('input');
    newInput.type = 'text';
    newInput.name = 'interests';
    newInput.placeholder = 'Nouveau centre d\'intérêt';
    interestsContainer.appendChild(newInput);
}

function updateProfile() {
    const name = document.getElementById('name').value;
    const bio = document.getElementById('bio').value;
    const age = document.getElementById('age-input').value;
    const location = document.getElementById('location-input').value;
    const email = document.getElementById('email-input').value;

    document.querySelector('.profile-name').innerText = name;
    document.querySelector('.profile-bio').innerText = bio;
    document.getElementById('age').innerText = age;
    document.getElementById('location').innerText = location;
    document.getElementById('email').innerText = email;


    const interestsInputs = document.querySelectorAll('#interests-inputs input');
    const interestsList = document.getElementById('interests-list');
    interestsList.innerHTML = ''; // Vider la liste actuelle

    interestsInputs.forEach(input => {
        if (input.value.trim() !== '') {
            const newInterest = document.createElement('li');
            newInterest.innerText = input.value;
            interestsList.appendChild(newInterest);
        }
    });

    alert('Profil mis à jour avec succès!');
}


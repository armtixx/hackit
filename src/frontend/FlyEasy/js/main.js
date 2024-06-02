// const navBtn = document.querySelector('.nav-icon-btn');
// const navIcon = document.querySelector('.nav-icon');
// const nav = document.querySelector('.header_top-row');

// navBtn.onclick = function () {
//     navIcon.classList.toggle('nav-icon--active');
//     nav.classList.toggle('header_top-row--mobile');
//     document.body.classList.toggle('no-scroll');
// }

const carousel = document.getElementById('carousel');
const scrollAmount = 200;

document.getElementById('scroll-left').addEventListener('click', () => {
    carousel.scrollBy({ left: -scrollAmount, behavior: 'smooth' });
});

document.getElementById('scroll-right').addEventListener('click', () => {
    carousel.scrollBy({ left: scrollAmount, behavior: 'smooth' });
});

// Меню
document.addEventListener('DOMContentLoaded', function () {
    const burgerMenu = document.getElementById('burger-menu');
    const navList = document.getElementById('nav-list');

    burgerMenu.addEventListener('click', function () {
        navList.classList.toggle('active');
        burgerMenu.classList.toggle('active');
    });
});

// дата в инпутах
document.addEventListener("DOMContentLoaded", function() {
    const dateInputs = document.querySelectorAll('data-input');
    
    dateInputs.forEach(input => {
        input.addEventListener("focus", function() {
            this.type = "date";
        });
        
        input.addEventListener("blur", function() {
            if (this.value === "") {
                this.type = "text";
            }
        });
    });
});

document.addEventListener("DOMContentLoaded", function() {
    const dateInputs = document.querySelectorAll('search-input');
    
    dateInputs.forEach(input => {
        input.addEventListener("focus", function() {
            this.type = "date";
        });
        
        input.addEventListener("blur", function() {
            if (this.value === "") {
                this.type = "text";
            }
        });
    });
});

// Добавление в избранное
document.addEventListener('DOMContentLoaded', () => {
    const favoriteIcons = document.querySelectorAll('.favorite-icon');
    
    favoriteIcons.forEach(icon => {
        icon.addEventListener('click', (e) => {
            const icon = e.target;
            toggleFavorite(icon);
        });
    });
});

function toggleFavorite(icon) {
    icon.classList.toggle('favorite');
}

// Регистрация обработка ошибок

document.getElementById('togglePassword').addEventListener('click', function() {
    const passwordField = document.getElementById('password');
    const type = passwordField.getAttribute('type') === 'password' ? 'text' : 'password';
    passwordField.setAttribute('type', type);

    this.textContent = type === 'password' ? '👁️' : '🙈';
});

document.getElementById('registrationForm').addEventListener('submit', function(event) {
    event.preventDefault();
    let valid = true;

    // Clear previous error messages
    document.querySelectorAll('.error-message').forEach(span => span.style.display = 'none');

    // Validate email
    const email = document.getElementById('email').value;
    if (!validateEmail(email)) {
        valid = false;
        document.getElementById('emailError').textContent = 'Введите корректный email.';
        document.getElementById('emailError').style.display = 'block';
    }

    // Validate username
    const username = document.getElementById('username').value;
    if (username.length < 3) {
        valid = false;
        document.getElementById('usernameError').textContent = 'Имя пользователя должно быть не менее 3 символов.';
        document.getElementById('usernameError').style.display = 'block';
    }

    // Validate password
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;
    if (password.length < 6) {
        valid = false;
        document.getElementById('passwordError').textContent = 'Пароль должен быть не менее 6 символов.';
        document.getElementById('passwordError').style.display = 'block';
    } else if (password !== confirmPassword) {
        valid = false;
        document.getElementById('confirmPasswordError').textContent = 'Пароли не совпадают.';
        document.getElementById('confirmPasswordError').style.display = 'block';
    }

    // If form is valid, submit the form
    if (valid) {
        this.submit();
    }
});

function validateEmail(email) {
    const re = /^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/;
    return re.test(String(email).toLowerCase());
}

// Кнопка показать ещё
document.querySelector('.more-places').addEventListener('click', function() {
    const hiddenDestinations = document.querySelectorAll('.destination.hidden');
    hiddenDestinations.forEach(destination => {
        destination.classList.remove('hidden');
    });

    if (hiddenDestinations.length === 0) {
        this.style.display = 'none'; // Скрыть кнопку после отображения всех скрытых элементов
    }
    this.style.display = 'none'; // Скрыть кнопку после нажатия на неё
});


document.querySelectorAll('.checkbox').forEach(button => {
    button.addEventListener('click', () => {
        document.querySelectorAll('.checkbox').forEach(btn => btn.classList.remove('active'));
        button.classList.add('active');
    });
});

flatpickr('#departure', {
    dateFormat: 'd-m-Y'
});

flatpickr('#return', {
    dateFormat: 'd-m-Y'
});

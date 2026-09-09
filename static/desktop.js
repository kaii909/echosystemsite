let currentZIndex = 2; // z inddex of the windows

function openWindow(title) {
}

function closeWindow(id) {
    document.getElementById(id).style.display = "none";
}

// z index logic
function bringToFront(windowElement) {
    currentZIndex++;
    windowElement.style.zIndex = currentZIndex;
}

document.addEventListener('DOMContentLoaded', () => {
    const windows = document.querySelectorAll('.window');

    windows.forEach(win => {
        const titleBar = win.querySelector('.title-bar');

        titleBar.addEventListener('mousedown', (e) => {
            bringToFront(win); // Traz para frente
            
            let shiftX = e.clientX - win.getBoundingClientRect().left;
            let shiftY = e.clientY - win.getBoundingClientRect().top;

            document.onmousemove = (e) => {
                let newX = e.clientX - shiftX;
                let newY = e.clientY - shiftY;

                win.style.left = newX + 'px';
                win.style.top = newY + 'px';
            };

            document.onmouseup = () => {
                document.onmousemove = null;
                document.onmouseup = null;
            };
        });
    });
});

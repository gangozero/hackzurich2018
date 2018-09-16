function drawTrain(train, gridSize, ctx, canvas) {
    let x = train.x;
    let y = train.y;
    let width = train.width;
    let height = train.height;

    drawTrainRect(0, y + 1, width + 2, 1, gridSize, "#3f3f3f", ctx);
    drawTrainRect(0, y + 4, width + 2, 1, gridSize, "#3f3f3f", ctx);

    drawTrainRect(x, y, width, height, gridSize, "#dcdce2", ctx);
    drawTrainRect(x + 1, y + 1, width - 2, height - 2, gridSize, "#3468e2", ctx);

    drawTrainRect(x + Math.floor(width / 3), y, 1, height, gridSize, "#dcdce2", ctx);
    drawTrainRect(x + Math.floor((width / 3) * 2), y, 1, height, gridSize, "#dcdce2", ctx);

    drawTrainRect(x, y + 1, 1, 1, gridSize, "#ff0001", ctx);
    drawTrainRect(x, y + 4, 1, 1, gridSize, "#ff0001", ctx);

    drawTrainRect(x + width - 1, y + 1, 1, 1, gridSize, "#fff4a8", ctx);
    drawTrainRect(x + width - 1, y + 4, 1, 1, gridSize, "#fff4a8", ctx);


}

function drawTrainRect(x, y, width, height, gridSize, color, ctx) {
    x *= gridSize;
    y *= gridSize;
    width *= gridSize;
    height *= gridSize;

    ctx.fillStyle = color;
    ctx.fillRect(x, y, width, height);
}


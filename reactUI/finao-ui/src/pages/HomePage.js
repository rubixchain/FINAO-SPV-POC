import React, { useEffect, useRef } from 'react';
import './HomePage.css';

function HomePage() {
  const canvasRef = useRef(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    const ctx = canvas.getContext('2d');

    function rand(min, max) {
      return Math.floor(Math.random() * (max - min + 1) + min);
    }

    // canvas 
    var X = canvas.width = window.innerWidth;
    var Y = canvas.height = window.innerHeight;
    var mouseX = X / 2;
    var mouseY = Y / 2;

    var maxRadius = 30;
    var flowerNumX = X / (maxRadius * 4) + 1;
    var flowerNumY = Y / (maxRadius * 4) + 1;
    var flowers = [];

    function Flower(ctx, x, y) {
      this.ctx = ctx;
      this.init(x, y);
    }

    Flower.prototype.init = function(x, y) {
      this.x = x;
      this.y = y;
      this.a = 0;
      this.rad = this.a * Math.PI / 180;
      this.dist = 0;
      this.r = maxRadius / 2;
      this.c = {
        r: rand(0, 255),
        g: rand(0, 255),
        b: rand(0, 255)
      };
      this.flg = false;
    };

    Flower.prototype.draw = function() {
      var ctx = this.ctx;
      ctx.save();
      ctx.globalCompositeOperation = 'lighter';
      ctx.fillStyle = 'rgb(' + this.c.r + ', ' + this.c.g + ', ' + this.c.b + ')';
      ctx.globalAlpha = 0.5;
      ctx.translate(this.x, this.y);
      ctx.rotate(this.rad);
      ctx.translate(-this.x, -this.y);
      for (var i = 0; i < 6; i++) {
        ctx.beginPath();
        ctx.arc(Math.cos(60 * Math.PI / 180 * i) * this.dist + this.x, Math.sin(60 * Math.PI / 180 * i) * this.dist + this.y, this.r, 0, Math.PI * 2, false);
        ctx.fill();
      }
      ctx.restore();
    };

    Flower.prototype.updateParams = function() {
      if (this.flg === false) {
        this.r += maxRadius * 0.01 * 0.25;
        this.dist += maxRadius * 0.005;
        this.a += maxRadius * 0.005;
      }
      if (this.flg === true) {
        this.r -= maxRadius * 0.01 * 0.25;
        this.dist -= maxRadius * 0.005;
        this.a -= maxRadius * 0.005;
      }
      if (this.dist < 0) {
        this.changeColor();
      }
      this.rad = this.a * Math.PI / 180;
    };

    Flower.prototype.changeColor = function() {
      this.c = {
        r: rand(0, 255),
        g: rand(0, 255),
        b: rand(0, 255)
      };
    };

    Flower.prototype.render = function() {
      this.updateParams();
      this.turning();
      this.draw();
    };

    Flower.prototype.turning = function() {
      if (this.dist > maxRadius) this.flg = true;
      if (this.dist < 0) this.flg = false; 
    };

    for (var i = 0; i < flowerNumX; i++) {
      for (var j = 0; j < flowerNumY; j++) {
        var flower = new Flower(ctx, maxRadius * 4 * i, maxRadius * 4 * j);
        flowers.push(flower);
      } 
    }

    function render(){
      ctx.clearRect(0, 0, X, Y);
      for (var i = 0; i < flowers.length; i ++) {
        flowers[i].render();
      }
      requestAnimationFrame(render);
    }

    // resize
    function onResize() {
      X = canvas.width = window.innerWidth;
      Y = canvas.height = window.innerHeight;
      mouseX = X / 2;
      mouseY = Y / 2;
      flowerNumX = X / (maxRadius * 4) + 1;
      flowerNumY = Y / (maxRadius * 4) + 1;
      flowers = [];
      for (var i = 0; i < flowerNumX; i++) {
        for (var j = 0; j < flowerNumY; j++) {
          var flower = new Flower(ctx, maxRadius * 4 * i, maxRadius * 4 * j);
          flowers.push(flower);
        } 
      }
    }

    window.addEventListener('load', function() {
      if (!canvas || !canvas.getContext) {
        return false;
      }

      render();
      window.addEventListener('resize', onResize);
    });

    // Cleanup function
    return () => {
      window.removeEventListener('resize', onResize);
    };
  }, []);

  return <canvas ref={canvasRef} id="canvas">Canvas not supported.</canvas>;
}

export default HomePage;

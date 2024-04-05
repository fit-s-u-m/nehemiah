import * as THREE from '../../static/node_modules/three/build/three.module.js';

const scene = new THREE.Scene(); // create a scene

const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
camera.position.set(0, 0, 100);
camera.lookAt(0, 0, 0);

const renderer = new THREE.WebGLRenderer();
renderer.setSize(window.innerWidth, window.innerHeight); // set canvas to the window width and height
document.body.appendChild(renderer.domElement);
renderer.setClearColor(0xfff6e6);


const length = 50

const points = [];
points.push(new THREE.Vector3(- length, 0, 0));
points.push(new THREE.Vector3(length, 0, 0));

const point2 = [];
point2.push(new THREE.Vector3(0, -length, 0));
point2.push(new THREE.Vector3(0, length, 0));

// ###############################

const material = new THREE.LineBasicMaterial({ color: 0x0000ff });
const material2 = new THREE.LineBasicMaterial({ color: 0xff0000 });

const geometry = new THREE.BufferGeometry().setFromPoints(points);
const geometry2 = new THREE.BufferGeometry().setFromPoints(point2);

const line = new THREE.Line(geometry, material);
const line2 = new THREE.Line(geometry2, material2);
scene.add(line);
scene.add(line2);
renderer.render(scene, camera);


// function animate() {
//   requestAnimationFrame(animate);
//
//   // cube.rotation.x += 0.01;
//   // cube.rotation.y += 0.01;
//
//   renderer.render(scene, camera);
// }
//
// animate();

import { createRouter, createWebHistory } from 'vue-router';
import SchoolList from '../components/schools/SchoolList.vue';
import EventList from '../components/events/EventList.vue';
import AttendeeList from '../components/attendees/AttendeeList.vue';
import Capture from '../components/attendances/Capture.vue';

const routes = [
  { path: '/attendees', component: AttendeeList },
  { path: '/attendance', component: Capture },
  { path: '/events', component: EventList },
  { path: '/schools', component: SchoolList },
  { path: '/', redirect: '/attendees' } // Default path
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

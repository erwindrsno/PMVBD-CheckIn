import { createRouter, createWebHistory } from 'vue-router';
import SchoolList from '../components/schools/SchoolList.vue';
import EventList from '../components/events/EventList.vue';
import AttendeeList from '../components/attendees/AttendeeList.vue';
import Capture from '../components/attendances/Capture.vue';
import AttendanceList from '../components/attendances/AttendanceList.vue';

const routes = [
  { path: '/attendees', component: AttendeeList },
  { path: '/attendances/capture', component: Capture },
  { path: '/attendances/list', component: AttendanceList },
  { path: '/events', component: EventList },
  { path: '/schools', component: SchoolList },
  { path: '/', redirect: '/attendees' } // Default path
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

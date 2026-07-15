<script setup>
import { ref } from 'vue';
import SchoolList from './components/schools/SchoolList.vue';
import EventList from './components/events/EventList.vue';
import AttendeeList from './components/attendees/AttendeeList.vue';
import Capture from './components/attendances/Capture.vue';

const currentView = ref('attendee');
</script>

<template>
  <v-app>
    <v-app-bar rounded>
      <v-app-bar-title>PMVBD, Check In!</v-app-bar-title>
      <template v-slot:append>
        <v-btn @click="currentView = 'attendee'">Attendee</v-btn>
        <v-btn @click="currentView = 'attendance'">Attendance</v-btn>
        <v-btn @click="currentView = 'event'">Event</v-btn>
        <v-btn @click="currentView = 'school'">School</v-btn>
      </template>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <!-- Clean Component Switching -->
        <SchoolList v-if="currentView === 'school'" />
        <EventList v-else-if="currentView === 'event'" />
        <AttendeeList v-else-if="currentView === 'attendee'" />
        <Capture v-else-if="currentView === 'attendance'" />
        <div v-else><h1>Select a module</h1></div>
      </v-container>
    </v-main>
  </v-app>
</template>

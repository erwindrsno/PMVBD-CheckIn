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
        <!-- Use 'to' instead of @click -->
        <v-btn to="/attendees">Attendee</v-btn>
        <v-btn to="/attendance">Attendance</v-btn>
        <v-btn to="/events">Event</v-btn>
        <v-btn to="/schools">School</v-btn>
      </template>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <!-- The component automatically swaps here based on the URL -->
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

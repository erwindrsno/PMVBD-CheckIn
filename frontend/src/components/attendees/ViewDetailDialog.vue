<script setup>
import QRCode from 'qrcode';

defineProps({
  modelValue: { type: Boolean, required: true },
  selectedAttendee: { type: Object, default: null }
});

defineEmits(['update:modelValue']); // No need to emit 'print-qr' anymore

const downloadQR = async (attendee) => {
  const dataToEncode = attendee.public_id;

  try {
    // 1. Generate QR
    const qrDataUrl = await QRCode.toDataURL(dataToEncode, {
      errorCorrectionLevel: 'H',
      width: 500,
      margin: 2
    });

    const [qrImage, logoImage] = await Promise.all([
      loadImage(qrDataUrl), // Ensure you have the loadImage helper function
      loadImage('../../../assets/LOGO_PMV.PNG')
    ]);

    const canvas = document.createElement('canvas');
    canvas.width = 500;
    canvas.height = 500;
    const ctx = canvas.getContext('2d');

    // 2. Draw the QR code
    ctx.drawImage(qrImage, 0, 0);

    // 3. Define Circular Area
    const logoSize = 120;
    const centerX = 500 / 2;
    const centerY = 500 / 2;
    const radius = logoSize / 2;

    // 4. Create circular cutout
    ctx.save();
    ctx.beginPath();
    ctx.arc(centerX, centerY, radius, 0, Math.PI * 2);
    ctx.clip(); // Creates a circular mask

    // Clear everything inside the circle
    ctx.clearRect(centerX - radius, centerY - radius, logoSize, logoSize);

    // Optional: Draw a white background inside the circle
    // to ensure the logo has a clean base against the QR dots
    ctx.fillStyle = 'white';
    ctx.fill();

    ctx.restore();

    // 5. Draw the Logo inside the circular area
    ctx.save();
    ctx.beginPath();
    ctx.arc(centerX, centerY, radius, 0, Math.PI * 2);
    ctx.clip(); // Clip the logo to the same circle
    ctx.drawImage(logoImage, centerX - radius, centerY - radius, logoSize, logoSize);
    ctx.restore();

    // 6. Trigger download
    const link = document.createElement('a');
    link.href = canvas.toDataURL('image/png');
    link.download = `QR_${attendee.name.replace(/\s+/g, '_')}.png`;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

  } catch (err) {
    console.error('Failed to generate QR with circle logo:', err);
  }
};
</script>

<template>
  <v-dialog :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)" max-width="600px">
    <v-card v-if="selectedAttendee">
      <v-card-title class="bg-primary text-white">Attendee Details</v-card-title>
      <v-card-text class="mt-4">
        <v-list density="compact">
          <v-list-item title="Name" :subtitle="selectedAttendee.name"></v-list-item>
          <v-list-item title="ID" :subtitle="selectedAttendee.public_id"></v-list-item>
          <v-list-item title="School" :subtitle="selectedAttendee.school"></v-list-item>
          <v-list-item title="Grade" :subtitle="`${selectedAttendee.grade} - ${selectedAttendee.subgrade}`"></v-list-item>
          <v-list-item title="Phone" :subtitle="selectedAttendee.telp"></v-list-item>
          <v-list-item title="Guardian Phone" :subtitle="selectedAttendee.guardian_telp"></v-list-item>
        </v-list>
      </v-card-text>

      <v-card-actions>
        <!-- New Print QR Button -->
        <v-btn
          variant="flat"
          color="secondary"
          prepend-icon="mdi-qrcode"
          @click="downloadQR(selectedAttendee)"
        >
          Print QR
        </v-btn>

        <v-spacer></v-spacer>
        <v-btn color="primary" @click="$emit('update:modelValue', false)">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

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

    // 2. Load Images
    const qrImage = new Image();
    qrImage.src = qrDataUrl;
    await qrImage.decode();

    const logoImage = new Image();
    logoImage.src = '../../../assets/LOGO_PMV.PNG';
    await logoImage.decode();

    // 3. Setup Canvas
    const canvas = document.createElement('canvas');
    canvas.width = 500;
    canvas.height = 500;
    const ctx = canvas.getContext('2d');

    // 4. Draw QR Code
    ctx.drawImage(qrImage, 0, 0);

    // 5. Logo settings
    const logoSize = 120;
    const center = canvas.width / 2;
    const x = center - logoSize / 2;
    const y = center - logoSize / 2;

    // Radius of the white circle (slightly larger than the logo)
    const circleRadius = logoSize / 2 + 2;

    // 6. Draw white circular background
    ctx.beginPath();
    ctx.arc(center, center, circleRadius, 0, Math.PI * 2);
    ctx.closePath();
    ctx.fillStyle = '#FFFFFF';
    ctx.fill();

    // Optional: Add a white border for a cleaner look
    ctx.lineWidth = 4;
    ctx.strokeStyle = '#FFFFFF';
    ctx.stroke();

    // 7. Clip logo into a circle
    ctx.save();

    ctx.beginPath();
    ctx.arc(center, center, logoSize / 2, 0, Math.PI * 2);
    ctx.closePath();
    ctx.clip();

    // Draw logo
    ctx.drawImage(logoImage, x, y, logoSize, logoSize);

    ctx.restore();

    // 8. Download
    const link = document.createElement('a');
    link.href = canvas.toDataURL('image/png');
    link.download = `QR_${attendee.name.replace(/\s+/g, '_')}.png`;

    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

  } catch (err) {
    console.error('Failed to generate QR with logo:', err);
  }
};
</script>

<template>
  <v-dialog :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)" max-width="600px">
    <v-card v-if="selectedAttendee">
      <v-card-title class="bg-primary text-white">Detail Peserta</v-card-title>
      <v-card-text class="mt-4">
        <v-list density="compact">
          <v-list-item title="Nama" :subtitle="selectedAttendee.name"></v-list-item>
          <v-list-item title="ID" :subtitle="selectedAttendee.public_id"></v-list-item>
          <v-list-item title="Sekolah" :subtitle="selectedAttendee.school"></v-list-item>
          <v-list-item title="Kelas" :subtitle="`${selectedAttendee.grade} - ${selectedAttendee.subgrade}`"></v-list-item>
          <v-list-item title="No. Telp" :subtitle="selectedAttendee.telp"></v-list-item>
          <v-list-item title="No. Telp Orang tua/wali" :subtitle="selectedAttendee.guardian_telp"></v-list-item>
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
          Cetak QR
        </v-btn>

        <v-spacer></v-spacer>
        <v-btn color="primary" @click="$emit('update:modelValue', false)">Tutup</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

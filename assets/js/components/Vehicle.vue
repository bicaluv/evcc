<template>
	<div>
		<div class="mb-2">{{ socTitle || "Fahrzeug" }} 
				<fa-icon
					v-if="minSoCNighttimeActive"
					icon="bed"
					:class="{
						'text-primary': isNighttime,
						'text-muted': !isNighttime,
					}"
				></fa-icon>
		</div>
		<div class="progress" style="height: 24px; font-size: 100%; margin-top: 16px">
			<div
				class="progress-bar"
				role="progressbar"
				:class="{
					'progress-bar-striped': charging,
					'progress-bar-animated': charging,
					'bg-light': !connected,
					'text-secondary': !connected,
					'bg-warning': connected && minSoCActive,
				}"
				:style="{ width: socChargeDisplayWidth + '%' }"
			>
				{{ socChargeDisplayValue }}
			</div>
			<div
				class="progress-bar"
				role="progressbar"
				:class="{
					'progress-bar-striped': charging,
					'progress-bar-animated': charging,
					'bg-warning': true,
					'bg-muted': true,
				}"
				:style="{ width: minSoCRemainingDisplayWidth + '%' }"
				v-if="minSoCActive && socChargeDisplayWidth < 100"
			></div>
		</div>
	</div>
</template>

<script>
import "../icons";

export default {
	name: "Vehicle",
	props: {
		socTitle: String,
		connected: Boolean,
		charging: Boolean,
		hasVehicle: Boolean,
		socCharge: Number,
		minSoC: Number,
		minSoCGeoLat: Number,
		minSoCGeoLong: Number,
		isNighttime: Boolean,
	},
	computed: {
		socChargeDisplayWidth: function () {
			if (this.hasVehicle && this.socCharge >= 0) {
				return this.socCharge;
			}
			return 100;
		},
		socChargeDisplayValue: function () {
			// no soc or no soc value
			if (!this.hasVehicle || this.socCharge < 0) {
				let chargeStatus = "getrennt";
				if (this.charging) {
					chargeStatus = "laden";
				} else if (this.connected) {
					chargeStatus = "verbunden";
				}
				return chargeStatus;
			}

			// percent value if enough space
			let socCharge = this.socCharge;
			if (socCharge >= 10) {
				socCharge += "%";
			}
			return socCharge;
		},
		minSoCActive: function () {
			return this.minSoC > 0 && this.socCharge < this.minSoC;
		},
		minSoCNighttimeActive: function () {
			return this.minSoCGeoLat > 0 && this.minSoCGeoLong > 0;
		},
		minSoCRemainingDisplayWidth: function () {
			return this.minSoC - this.socCharge;
		},
	},
};
</script>

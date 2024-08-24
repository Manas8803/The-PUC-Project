import type { Config } from "tailwindcss";

const config = {
	darkMode: ["class"],
	content: [
		"./pages/**/*.{ts,tsx}",
		"./components/**/*.{ts,tsx}",
		"./app/**/*.{ts,tsx}",
		"./src/**/*.{ts,tsx}",
	],
	prefix: "",
	theme: {
		colors: {
			bgrnd: "#F5F5F5",
			main: "#7CC558",
			gr_white: "#FCEBED",
			gr_red: "#FF4848",
			side: "#4e7a37",
			white: "#FFFFFF",
			black: "#000000",
			test_1: "#00FF00",
			test_2: "#FF0000",
			test_3: "#0000FF",
			gray_50: "#F9FAFB",
			gray_100: "#F3F4F6",
			gray_200: "#E5E7EB",
			gray_300: "#D1D5DB",
			gray_400: "#9CA3AF",
			gray_500: "#6B7280",
			gray_600: "#4B5563",
			gray_700: "#374151",
			gray_800: "#1F2937",
			gray_900: "#111827",
			gray_950: "#030712",
			green_50: "#F0FDF4",
			green_100: "#DCFCE7",
			green_200: "#BBF7D0",
			green_300: "#86EFAC",
			green_400: "#4ADE80",
			green_500: "#22C55E",
			green_600: "#16A34A",
			green_700: "#15803D",
			green_800: "#166534",
			green_900: "#14532D",
			green_950: "#052E16",
		},
		container: {
			center: true,
			padding: "2rem",
			screens: {
				"2xl": "1400px",
			},
		},

		extend: {
			transitionDelay: {
				"0": "0ms",
				"400": "400ms",
				"600": "600ms",
				"2000": "2000ms",
				"5000": "5000ms",
			},
			transitionProperty: {
				height: "height",
				display: "display",
			},
			transitionDuration: {
				"0": "0ms",
				"400": "400ms",
				"600": "600ms",
				"2000": "2000ms",
				"5000": "5000ms",
			},
			colors: {
				border: "hsl(var(--border))",
				input: "hsl(var(--input))",
				ring: "hsl(var(--ring))",
				background: "hsl(var(--background))",
				foreground: "hsl(var(--foreground))",
				primary: {
					DEFAULT: "hsl(var(--primary))",
					foreground: "hsl(var(--primary-foreground))",
				},
				secondary: {
					DEFAULT: "hsl(var(--secondary))",
					foreground: "hsl(var(--secondary-foreground))",
				},
				destructive: {
					DEFAULT: "hsl(var(--destructive))",
					foreground: "hsl(var(--destructive-foreground))",
				},
				muted: {
					DEFAULT: "hsl(var(--muted))",
					foreground: "hsl(var(--muted-foreground))",
				},
				accent: {
					DEFAULT: "hsl(var(--accent))",
					foreground: "hsl(var(--accent-foreground))",
				},
				popover: {
					DEFAULT: "hsl(var(--popover))",
					foreground: "hsl(var(--popover-foreground))",
				},
				card: {
					DEFAULT: "hsl(var(--card))",
					foreground: "hsl(var(--card-foreground))",
				},
			},
			borderRadius: {
				lg: "var(--radius)",
				md: "calc(var(--radius) - 2px)",
				sm: "calc(var(--radius) - 4px)",
			},
			keyframes: {
				"accordion-down": {
					from: { height: "0" },
					to: { height: "var(--radix-accordion-content-height)" },
				},
				"accordion-up": {
					from: { height: "var(--radix-accordion-content-height)" },
					to: { height: "0" },
				},
			},
			animation: {
				"accordion-down": "accordion-down 0.2s ease-out",
				"accordion-up": "accordion-up 0.2s ease-out",
			},
		},
	},
	plugins: [require("tailwindcss-animate")],
} satisfies Config;

export default config;

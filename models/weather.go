package models

type GameMoodData struct {
	Theme        string
	ColorPalette string
	Environment  GameEnvironment
	stats        PlayerModifiers
}

type GameEnvironment struct {
	SkyboxTexture         string
	FogDensity            float32
	WeatherParticleEffect string
}

type PlayerModifiers struct {
	StaminaDrainRate float32
	MovementSpeed    float32
}

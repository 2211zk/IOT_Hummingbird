module.exports = {
  displayName: 'StandardCategorySelector',
  testMatch: [
    '<rootDir>/src/components/standardCategorySelector/__tests__/**/*.test.js'
  ],
  moduleFileExtensions: ['js', 'json', 'vue'],
  transform: {
    '^.+\\.vue$': '@vue/vue3-jest',
    '^.+\\.js$': 'babel-jest'
  },
  moduleNameMapping: {
    '^@/(.*)$': '<rootDir>/src/$1'
  },
  testEnvironment: 'jsdom',
  setupFilesAfterEnv: ['<rootDir>/src/components/standardCategorySelector/__tests__/setup.js'],
  collectCoverageFrom: [
    'src/components/standardCategorySelector/**/*.{js,vue}',
    '!src/components/standardCategorySelector/__tests__/**',
    '!src/components/standardCategorySelector/utils/**'
  ],
  coverageReporters: ['text', 'lcov', 'html'],
  coverageDirectory: '<rootDir>/coverage/standardCategorySelector',
  clearMocks: true,
  restoreMocks: true
}
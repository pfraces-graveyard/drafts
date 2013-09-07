# Type

Type system inspired in go interfaces

# Built-in types

*   `null`
*   `boolean`
*   `number`
*   `string`
*   `array`
*   `function`

# New type definition

    var type = require('type');

    type.new('point', {
      x: 'number',
      y: 'number'
    });

# Type checking

    expect(type.is('null', null)).toBe(true);
    expect(type.is('null', undefined)).toBe(true);
    expect(type.is('null')).toBe(true);
    expect(type.is('null', 0)).toBe(false);
    expect(type.is('null', false)).toBe(false);
    expect(type.is('null', '')).toBe(false);

    expect(type.is('boolean', true)).toBe(true);
    expect(type.is('boolean', false)).toBe(true);
    expect(type.is('boolean', 'foo')).toBe(false);
    expect(type.is('boolean', 1)).toBe(false);
    expect(type.is('boolean', 0)).toBe(false);
    expect(type.is('boolean', null)).toBe(false);
    
    expect(type.is('number', 0)).toBe(true);
    expect(type.is('number', 1)).toBe(true);
    expect(type.is('number', 'foo')).toBe(false);
    expect(type.is('number', true)).toBe(false);

    expect(type.is('string', 'foo')).toBe(true);
    expect(type.is('string', true)).toBe(false);
    expect(type.is('string', 0)).toBe(false);
    expect(type.is('string', '')).toBe(true);
    expect(type.is('string', '')).toBe(true);

    expect(type.is('point', { x: 0, y: 0 })).toBe(true);
    expect(type.is('point', true).toBe(false);
    expect(type.is('point', { x: 0 })).toBe(false);

# Type aliases

    type.alias({ boolean: 'bool' });

    type.alias({
      boolean: 'bool',
      string: ['str', 's']
    });

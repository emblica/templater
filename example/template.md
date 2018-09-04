

# {{ title }}

Is a novel authored by {{ author }}

### List of some characters in the novel:

{% for character in characters %}
  - {{ character.name }} _({{ character.type }})_
{% endfor %}
